from django.db import models

from order.models import Order
from product.models import Product


# Create your models here.
class Review(models.Model):
    comment = models.TextField(blank=True, null=True)
    rating = models.PositiveIntegerField()
    positive_feedbacks_count = models.PositiveIntegerField(default=0)
    negative_feedbacks_count = models.PositiveIntegerField(default=0)
    abusive_reports_count = models.PositiveIntegerField(default=0)

    order = models.ForeignKey(Order, blank=False, null=False, on_delete=models.CASCADE)
    product = models.ForeignKey(Product, blank=True, null=True, on_delete=models.CASCADE)
    # product_variant = models.ForeignKey(ProductVariation, blank=True, null=True, on_delete=models.CASCADE)

    def __str__(self):
        return f"Order-{self.order}"
