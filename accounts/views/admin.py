from rest_framework import status
from rest_framework.response import Response

from accounts.models import User
from accounts.serializers import AdminListSerializer
from core.permissions.base import ROLE
from core.views.base import BaseViewSet


# Create your views here.
class AdminViewSet(BaseViewSet):
    model = User
    serializer_class = AdminListSerializer

    search_fields = ["username", "email"]

    def get_queryset(self):
        return (
            self.filter_queryset(super().get_queryset().select_related("type"))
        )

    def list(self, request, *args, **kwargs):
        admins = User.objects.filter(role=ROLE.ADMIN.value)
        serializer = self.serializer_class(admins, many=True)
        return Response(serializer.data, status=status.HTTP_200_OK)
