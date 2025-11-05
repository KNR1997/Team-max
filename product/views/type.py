from rest_framework import status
from rest_framework.response import Response

from core.permissions.base import allow_permission, ROLE
from core.views.base import BaseViewSet
from product.models import Type
from product.serializers import TypeSerializer


# Create your views here.
class TypeViewSet(BaseViewSet):
    model = Type
    serializer_class = TypeSerializer

    search_fields = ["name", "slug"]
    filterset_fields = []

    lookup_field = "slug"

    def get_queryset(self):
        return (
            self.filter_queryset(super().get_queryset())
        )

    def list(self, request, *args, **kwargs):
        types = Type.objects.all()
        serializer = TypeSerializer(types, many=True)
        return Response(serializer.data, status=status.HTTP_200_OK)

    def retrieve(self, request, *args, **kwargs):
        return super().retrieve(request, *args, **kwargs)

    @allow_permission([ROLE.ADMIN])
    def create(self, request, *args, **kwargs):
        return super().create(request, *args, **kwargs)

    @allow_permission([ROLE.ADMIN], creator=True, model=Type)
    def partial_update(self, request, *args, **kwargs):
        return super().partial_update(request, *args, **kwargs)

    @allow_permission([ROLE.ADMIN], creator=True, model=Type)
    def destroy(self, request, *args, **kwargs):
        return super().destroy(request, *args, **kwargs)
