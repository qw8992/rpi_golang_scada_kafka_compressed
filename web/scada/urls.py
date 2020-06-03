from django.urls import path
from . import views


app_name = 'scada'
urlpatterns = [
    path('', views.ResourceMonitoringView.as_view(), name='resource'),
]