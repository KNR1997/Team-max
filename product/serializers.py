from rest_framework import serializers
from slugify import slugify

from core.serializers.base import BaseSerializer
from .models import Type, Category, Tag, Product, Manufacturer, Author


class TypeSerializer(serializers.ModelSerializer):
    class Meta:
        model = Type
        fields = '__all__'


class CategorySerializer(BaseSerializer):
    class Meta:
        model = Category
        fields = '__all__'
        read_only_fields = ['slug']

    def create(self, validated_data):
        name = validated_data.get("name")
        slug = slugify(name)

        return Category.objects.create(
            **validated_data,
            slug=slug
        )


class CategoryListSerializer(BaseSerializer):
    type = TypeSerializer(read_only=True)

    class Meta:
        model = Category
        fields = '__all__'


class TagSerializer(serializers.ModelSerializer):
    class Meta:
        model = Tag
        fields = '__all__'
        read_only_fields = ['slug']

    def create(self, validated_data):
        name = validated_data.get("name")
        slug = slugify(name)

        return Tag.objects.create(
            **validated_data,
            slug=slug
        )


class TagListSerializer(serializers.ModelSerializer):
    type = TypeSerializer(read_only=True)

    class Meta:
        model = Tag
        fields = '__all__'


class ProductSerializer(serializers.ModelSerializer):
    class Meta:
        model = Product
        fields = '__all__'
        read_only_fields = ['slug', 'categories', 'tags']

    def create(self, validated_data):
        name = validated_data.get("name")
        slug = slugify(name)

        product = Product.objects.create(
            **validated_data, slug=slug
        )
        return product


class ProductListSerializer(BaseSerializer):
    type = TypeSerializer(read_only=True)
    categories = CategorySerializer(read_only=True, many=True)
    tags = TagSerializer(read_only=True, many=True)

    class Meta:
        model = Product
        fields = "__all__"


class ManufacturerSerializer(BaseSerializer):
    class Meta:
        model = Manufacturer
        fields = '__all__'


class ManufacturerListSerializer(BaseSerializer):
    class Meta:
        model = Manufacturer
        fields = '__all__'


class AuthorSerializer(BaseSerializer):
    class Meta:
        model = Author
        fields = '__all__'


class AuthorListSerializer(BaseSerializer):
    class Meta:
        model = Author
        fields = '__all__'
