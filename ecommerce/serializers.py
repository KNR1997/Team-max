from core.serializers.base import BaseSerializer
from .models import Tax, Shipping


class TaxSerializer(BaseSerializer):
    class Meta:
        model = Tax
        fields = '__all__'


class TaxListSerializer(BaseSerializer):
    class Meta:
        model = Tax
        fields = '__all__'


class ShippingSerializer(BaseSerializer):
    class Meta:
        model = Shipping
        fields = '__all__'


class ShippingListSerializer(BaseSerializer):
    class Meta:
        model = Shipping
        fields = '__all__'
