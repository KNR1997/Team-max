from django.urls import path, include
from rest_framework.routers import DefaultRouter

from shops.views import ShopViewSet

router = DefaultRouter()

urlpatterns = [
    path('', include(router.urls)),
    path(
        "shops/",
        ShopViewSet.as_view({"get": "list", "post": "create"}),
        name="shop",
    ),
    path(
        "shops/<str:slug>/",
        ShopViewSet.as_view({
            "get": "retrieve",
            "put": "update",
            "patch": "partial_update",
            "delete": "destroy",
        }),
        name="shop",
    ),
]
