from core.serializers.base import BaseSerializer
from .models import Coupon, FlashSale


class CouponSerializer(BaseSerializer):
    class Meta:
        model = Coupon
        fields = '__all__'


class CouponListSerializer(BaseSerializer):
    class Meta:
        model = Coupon
        fields = '__all__'


class FlashSaleSerializer(BaseSerializer):
    class Meta:
        model = FlashSale
        fields = '__all__'


class FlashSaleListSerializer(BaseSerializer):
    class Meta:
        model = FlashSale
        fields = '__all__'
