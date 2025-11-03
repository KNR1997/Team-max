from core.serializers.base import BaseSerializer
from .models import TermsAndConditions, FAQ


class FAQSerializer(BaseSerializer):
    class Meta:
        model = FAQ
        fields = '__all__'


class FAQListSerializer(BaseSerializer):
    class Meta:
        model = FAQ
        fields = '__all__'


class TermsAndConditionsSerializer(BaseSerializer):
    class Meta:
        model = TermsAndConditions
        fields = '__all__'


class TermsAndConditionsListSerializer(BaseSerializer):
    class Meta:
        model = TermsAndConditions
        fields = '__all__'
