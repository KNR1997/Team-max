from rest_framework import status
from rest_framework.response import Response

from core.views.base import BaseViewSet
from product.models import Author
from product.serializers import AuthorListSerializer, AuthorSerializer


# Create your views here.
class AuthorViewSet(BaseViewSet):
    model = Author
    serializer_class = AuthorListSerializer

    search_fields = ["name", "slug"]
    filterset_fields = []

    lookup_field = "slug"

    def get_queryset(self):
        return (
            self.filter_queryset(super().get_queryset())
        )

    def create(self, request, *args, **kwargs):
        serializer = AuthorSerializer(data=request.data)

        if serializer.is_valid(raise_exception=True):
            serializer.save()
            data = serializer.data
            return Response(data, status=status.HTTP_201_CREATED)

    def list(self, request, *args, **kwargs):
        return super().list(request, *args, **kwargs)

    def partial_update(self, request, *args, **kwargs):
        return super().partial_update(request, *args, **kwargs)

    def destroy(self, request, *args, **kwargs):
        return super().destroy(request, *args, **kwargs)
