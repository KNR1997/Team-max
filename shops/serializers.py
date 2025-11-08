from crum import get_current_user
from rest_framework import serializers
from slugify import slugify

from core.serializers.base import BaseSerializer
from shops.models import Shop


class ShopSerializer(serializers.ModelSerializer):
    class Meta:
        model = Shop
        fields = '__all__'
        read_only_fields = ['slug', 'owner']

    def create(self, validated_data):
        name = validated_data.get("name")
        slug = slugify(name)

        current_user = get_current_user()

        return Shop.objects.create(
            **validated_data,
            slug=slug,
            owner=current_user
        )


class ShopListSerializer(serializers.ModelSerializer):
    class Meta:
        model = Shop
        fields = '__all__'


class ShopLiteSerializer(BaseSerializer):
    class Meta:
        model = Shop
        fields = [
            "id",
            "name",
            "slug"
        ]
