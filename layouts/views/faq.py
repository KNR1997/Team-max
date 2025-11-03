from core.views.base import BaseViewSet
from layouts.models import FAQ
from layouts.serializers import FAQListSerializer


# Create your views here.
class FaqViewSet(BaseViewSet):
    model = FAQ
    serializer_class = FAQListSerializer

    search_fields = ["name", "slug"]
    filterset_fields = []

    def get_queryset(self):
        return (
            self.filter_queryset(super().get_queryset())
        )

    def create(self, request, *args, **kwargs):
        return super().create(request, *args, **kwargs)

    def list(self, request, *args, **kwargs):
        return super().list(request, *args, **kwargs)

    def partial_update(self, request, *args, **kwargs):
        return super().partial_update(request, *args, **kwargs)

    def destroy(self, request, *args, **kwargs):
        return super().destroy(request, *args, **kwargs)
