from django.db import models

from accounts.models import User
from core.models.base import BaseModel


# Create your models here.
class Shop(BaseModel):
    name = models.CharField(max_length=20, unique=True)
    slug = models.SlugField(max_length=20, unique=True)
    description = models.TextField(null=True, blank=True)
    cover_image = models.JSONField(null=True, blank=True)
    logo = models.JSONField(null=True, blank=True)
    is_active = models.BooleanField(default=True)
    address = models.JSONField(null=True, blank=True)
    settings = models.JSONField(null=True, blank=True)

    # --- Relationships ---
    owner = models.ForeignKey(User, related_name="shops", on_delete=models.CASCADE)

    class Meta:
        db_table = "shop"

    def __str__(self):
        return self.name
