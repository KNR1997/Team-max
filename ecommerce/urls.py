from django.urls import path
from rest_framework.routers import DefaultRouter

from ecommerce.views.shipping import ShippingViewSet
from ecommerce.views.tax import TaxViewSet
# from promotions.views.coupon import CouponViewSet
from promotions.views.flash_sale import FlashSaleViewSet

router = DefaultRouter()

urlpatterns = [
    path(
        "taxes/",
        TaxViewSet.as_view({"get": "list", "post": "create"}),
        name="taxes",
    ),
    path(
        "taxes/<uuid:pk>",
        TaxViewSet.as_view({
            "get": "retrieve",
            "put": "update",
            "patch": "partial_update",
            "delete": "destroy",
        }),
        name="taxes",
    ),
    path(
        "shippings/",
        ShippingViewSet.as_view({"get": "list", "post": "create"}),
        name="shipping",
    ),
    path(
        "shippings/<uuid:pk>",
        ShippingViewSet.as_view({
            "get": "retrieve",
            "put": "update",
            "patch": "partial_update",
            "delete": "destroy",
        }),
        name="shipping",
    ),

]
