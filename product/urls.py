from django.urls import path, include
from rest_framework.routers import DefaultRouter

from product.views.author import AuthorViewSet
from product.views.category import CategoryViewSet
from product.views.manufacturer import ManufacturerViewSet
from product.views.product import ProductViewSet
from product.views.tag import TagViewSet
from product.views.type import TypeViewSet

router = DefaultRouter()

urlpatterns = [
    path('', include(router.urls)),
    path(
        "types/",
        TypeViewSet.as_view({"get": "list", "post": "create"}),
        name="type",
    ),
    path(
        "types/<str:slug>/",
        TypeViewSet.as_view({
            "get": "retrieve",
            "put": "update",
            "patch": "partial_update",
            "delete": "destroy",
        }),
        name="type",
    ),
    path(
        "categories/",
        CategoryViewSet.as_view({"get": "list", "post": "create"}),
        name="category",
    ),
    path(
        "categories/<str:slug>/",
        CategoryViewSet.as_view({
            "get": "retrieve",
            "put": "update",
            "patch": "partial_update",
            "delete": "destroy",
        }),
        name="category",
    ),
    path(
        "tags/",
        TagViewSet.as_view({"get": "list", "post": "create"}),
        name="tag",
    ),
    path(
        "tags/<str:slug>/",
        TagViewSet.as_view({
            "get": "retrieve",
            "put": "update",
            "patch": "partial_update",
            "delete": "destroy",
        }),
        name="tag",
    ),
    path(
        "products/",
        ProductViewSet.as_view({"get": "list", "post": "create"}),
        name="product",
    ),
    path(
        "products/<int:pk>",
        ProductViewSet.as_view({
            "get": "retrieve",
            "put": "update",
            "patch": "partial_update",
            "delete": "destroy",
        }),
        name="product",
    ),
    path(
        "manufacturers/",
        ManufacturerViewSet.as_view({"get": "list", "post": "create"}),
        name="manufacturers",
    ),
    path(
        "manufacturers/<int:pk>",
        ManufacturerViewSet.as_view({
            "get": "retrieve",
            "put": "update",
            "patch": "partial_update",
            "delete": "destroy",
        }),
        name="manufacturers",
    ),
    path(
        "authors/",
        AuthorViewSet.as_view({"get": "list", "post": "create"}),
        name="authors",
    ),
    path(
        "authors/<int:pk>",
        AuthorViewSet.as_view({
            "get": "retrieve",
            "put": "update",
            "patch": "partial_update",
            "delete": "destroy",
        }),
        name="authors",
    ),
]
