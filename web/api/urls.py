from django.urls import path
from . import views

app_name = 'api'
urlpatterns = [
    path('get/resource/', views.apiPcMonitoring, name='resource'),
]