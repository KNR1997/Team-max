from django.db import IntegrityError
from rest_framework import status
from rest_framework.response import Response

from accounts.models import User
from accounts.serializers import AdminListSerializer, UserSerializer
from core.views.base import BaseViewSet


# Create your views here.
class UserViewSet(BaseViewSet):
    model = User
    serializer_class = AdminListSerializer

    search_fields = ["username", "email"]

    def get_queryset(self):
        return (
            self.filter_queryset(super().get_queryset().select_related("type"))
        )

    def list(self, request, *args, **kwargs):
        admins = User.objects.all()
        serializer = self.serializer_class(admins, many=True)
        return Response(serializer.data, status=status.HTTP_200_OK)

    def create(self, request):
        try:
            serializer = UserSerializer(data=request.data)

            if serializer.is_valid(raise_exception=True):
                serializer.save()
                data = serializer.data
                return Response(data, status=status.HTTP_201_CREATED)
        except IntegrityError as e:
            if "already exists" in str(e):
                return Response(
                    {"slug": "The user with the username already exists"},
                    status=status.HTTP_409_CONFLICT,
                )
