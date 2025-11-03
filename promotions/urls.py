from django.urls import path
from rest_framework.routers import DefaultRouter

from promotions.views.coupon import CouponViewSet
from promotions.views.flash_sale import FlashSaleViewSet

router = DefaultRouter()

urlpatterns = [
    path(
        "coupons/",
        CouponViewSet.as_view({"get": "list", "post": "create"}),
        name="coupons",
    ),
    path(
        "coupons/<str:code>",
        CouponViewSet.as_view({
            "get": "retrieve",
            "put": "update",
            "patch": "partial_update",
            "delete": "destroy",
        }),
        name="coupons",
    ),
    path(
        "flash-sale/",
        FlashSaleViewSet.as_view({"get": "list", "post": "create"}),
        name="flash-sale",
    ),
    path(
        "flash-sale/<str:slug>",
        FlashSaleViewSet.as_view({
            "get": "retrieve",
            "put": "update",
            "patch": "partial_update",
            "delete": "destroy",
        }),
        name="flash-sale",
    )

]
