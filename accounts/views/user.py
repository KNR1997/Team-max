from django.db import IntegrityError
from rest_framework import status
from rest_framework.response import Response

from accounts.models import User
from accounts.serializers import UserSerializer, UserListSerializer
from core.permissions.base import allow_permission, ROLE
from core.views.base import BaseViewSet, BaseAPIView


# Create your views here.
class UserViewSet(BaseViewSet):
    model = User
    serializer_class = UserListSerializer

    search_fields = ["username", "email"]

    def get_queryset(self):
        return (
            self.filter_queryset(super().get_queryset())
        )

    def list(self, request, *args, **kwargs):
        return super().list(request, *args, **kwargs)

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


class UserBlockEndpoint(BaseAPIView):
    @allow_permission([ROLE.ADMIN])
    def post(self, request):
        user_id = request.data.get("id")
        if not user_id:
            return Response(
                {"error": "id is required"}, status=status.HTTP_400_BAD_REQUEST
            )

        user = User.objects.get(id=user_id)

        user.is_active = False
        user.save()
        return Response(status=status.HTTP_204_NO_CONTENT)


class UserUnblockEndpoint(BaseAPIView):
    @allow_permission([ROLE.ADMIN])
    def post(self, request):
        user_id = request.data.get("id")
        if not user_id:
            return Response(
                {"error": "id is required"}, status=status.HTTP_400_BAD_REQUEST
            )

        user = User.objects.get(id=user_id)

        user.is_active = True
        user.save()
        return Response(status=status.HTTP_204_NO_CONTENT)
