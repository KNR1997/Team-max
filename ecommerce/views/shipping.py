# Create your views here.
from rest_framework import status
from rest_framework.response import Response

from core.permissions.base import allow_permission, ROLE
from core.views.base import BaseViewSet
from ecommerce.models import Shipping
from ecommerce.serializers import ShippingSerializer, ShippingListSerializer
from product.models import Type


# Create your views here.
class ShippingViewSet(BaseViewSet):
    model = Shipping
    serializer_class = ShippingSerializer

    search_fields = ["slug"]
    filterset_fields = []

    lookup_field = "slug"

    def get_queryset(self):
        return (
            self.filter_queryset(super().get_queryset())
        )

    def list(self, request, *args, **kwargs):
        shipping_list = Shipping.objects.all()
        serializer = ShippingListSerializer(shipping_list, many=True)
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
