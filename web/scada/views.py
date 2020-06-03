from django.views.generic import ListView, TemplateView
from scada.models import Device


class ResourceMonitoringView(TemplateView):
    template_name = 'scada/resource_monitoring.html'