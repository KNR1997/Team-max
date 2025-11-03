from django.db import models

from core.utils.default_language import default_translated_languages


# Create your models here.
class CouponType(models.TextChoices):
    FIXED = "fixed", "Fixed"
    PERCENTAGE = "percentage", "Percentage"


class Coupon(models.Model):
    code = models.CharField(max_length=20, unique=True)
    description = models.TextField(null=True, blank=True)
    type = models.CharField(max_length=20, choices=CouponType.choices, db_index=True)
    amount = models.FloatField()
    minimum_cart_amount = models.FloatField()
    active_from = models.DateTimeField()
    expire_at = models.DateTimeField()
    is_valid = models.BooleanField(default=False)
    # target = models.IntegerField()
    language = models.CharField(max_length=10)
    translated_languages = models.JSONField(default=default_translated_languages, blank=True)

    class Meta:
        db_table = "coupon"

    def __str__(self):
        return self.code


class FlashSale(models.Model):
    SALE_TYPE = [
        ('fixed', 'Fixed'),
        ('percentage', 'Percentage')
    ]
    title = models.TextField()
    slug = models.SlugField(unique=True)
    description = models.TextField(blank=True, null=True)
    start_date = models.DateTimeField()
    end_date = models.DateTimeField()
    type = models.CharField(max_length=50, choices=SALE_TYPE)  # 'fixed', 'percentage', etc.
    image = models.JSONField(default=dict, blank=True, null=True)
    cover_image = models.JSONField(default=dict, blank=True, null=True)
    language = models.CharField(max_length=10, default="en", blank=True)
    translated_languages = models.JSONField(default=default_translated_languages, blank=True)

    def __str__(self):
        return self.title
