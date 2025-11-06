from django.urls import path
from rest_framework_simplejwt.views import TokenRefreshView

from accounts.views.admin import AdminViewSet
from accounts.views.auth import SignupView, SigninView, CreateAdminView, MeView
from accounts.views.user import UserViewSet, UserBlockEndpoint, UserUnblockEndpoint

# from accounts.views import SignupView, SigninView

# from .views import RegisterView, CustomTokenObtainPairView

urlpatterns = [
    # path('register/', RegisterView.as_view(), name='register'),
    # # path('login/', TokenObtainPairView.as_view(), name='token_obtain_pair'),
    # path('signin/', CustomTokenObtainPairView.as_view(), name='token_obtain_pair'),
    path('token/refresh/', TokenRefreshView.as_view(), name='token_refresh'),
    path('register/', SignupView.as_view(), name='signup'),
    path('token', SigninView.as_view(), name='token'),
    path('create-admin/', CreateAdminView.as_view(), name='create_admin'),
    path('me', MeView.as_view(), name='me'),

    path(
        "admin/list/",
        AdminViewSet.as_view({"get": "list"}),
        name="admin",
    ),
    path(
        "users/",
        UserViewSet.as_view({"get": "list"}),
        name="user",
    ),
    path(
        "users/block-user",
        UserBlockEndpoint.as_view(),
        name="user-block",
    ),
    path(
        "users/unblock-user",
        UserUnblockEndpoint.as_view(),
        name="user-unblock",
    ),
]
