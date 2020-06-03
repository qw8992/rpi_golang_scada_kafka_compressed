from django.views.decorators.csrf import csrf_exempt
from django.http import Http404
from django.db import connection
from django.http import JsonResponse
from django.conf import settings

from datetime import datetime

import psutil
import json
import re


@csrf_exempt
def apiPcMonitoring(request):
    if request.method == 'GET':
        return JsonResponse({
            "nowTime": [
                str(datetime.now())
            ],
            "cpu": [
                psutil.cpu_percent()
            ],
            "memory": [
                psutil.virtual_memory().percent
            ],
            "disk": [
                psutil.disk_usage('/').percent
            ],
            "info": [
                "추후 추가 예정"
            ]
        }, json_dumps_params={'ensure_ascii': False})
    else:
        raise Http404("404 Not Found")
