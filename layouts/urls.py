from django.urls import path
from rest_framework.routers import DefaultRouter

from layouts.views.faq import FaqViewSet
from layouts.views.terms_and_conditions import TermsAndConditionsViewSet

router = DefaultRouter()

urlpatterns = [
    path(
        "faqs/",
        FaqViewSet.as_view({"get": "list", "post": "create"}),
        name="faqs",
    ),
    path(
        "faqs/<uuid:pk>",
        FaqViewSet.as_view({
            "get": "retrieve",
            "put": "update",
            "patch": "partial_update",
            "delete": "destroy",
        }),
        name="faqs",
    ),
    path(
        "terms-and-conditions/",
        TermsAndConditionsViewSet.as_view({"get": "list", "post": "create"}),
        name="terms-and-conditions",
    ),
    path(
        "terms-and-conditions/<uuid:pk>",
        TermsAndConditionsViewSet.as_view({
            "get": "retrieve",
            "put": "update",
            "patch": "partial_update",
            "delete": "destroy",
        }),
        name="terms-and-conditions",
    ),

]
