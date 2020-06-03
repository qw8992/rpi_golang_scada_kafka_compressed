from django.db import models
from django.core.validators import MaxValueValidator, MinValueValidator
from django.conf import settings


class Device(models.Model):
    mac_id = models.CharField(max_length=15, verbose_name='MAC ID', db_column='MAC_ID')
    name = models.CharField(max_length=50, db_column='NAME')
    host = models.CharField(max_length=15, db_column='HOST')
    port = models.IntegerField(default=502, db_column='PORT')
    UNIT_CHOICES = (
        (1, 'UYEG DEVICE'),
        (0, 'EXPORT DEVICE'),
    )
    unit_id = models.IntegerField(default=1, choices=UNIT_CHOICES, db_column='UNIT_ID')
    REMAP_CHOICES = (
        (1, 'REMAP v1'),
        (2, 'REMAP v2'),
    )
    remap_version = models.IntegerField(default=1, choices=REMAP_CHOICES, db_column='REMAP_VERSION')    
    process_interval = models.IntegerField(default=100, validators=[MaxValueValidator(1000), MinValueValidator(50)], db_column='PROCESS_INTERVAL')
    retry_cycle = models.IntegerField(default=100, db_column='RETRY_CYCLE')
    retry_count = models.IntegerField(default=3000, db_column='RETRY_COUNT')
    retry_conn_failed_count = models.IntegerField(default=100, db_column='RETRY_CONN_FAILED_COUNT')
    enabled = models.BooleanField(default=True, db_column='ENABLED')
    enabled.boolean = True
    enabled.short_description = 'Is enabled'

    class Meta:
        db_table = 'DEVICE'
        

class Log(models.Model):
    fk_id = models.IntegerField(db_column="FK_ID")
    log = models.TextField(db_column="LOG")
    create_at = models.DateTimeField(auto_now_add=True, db_column="CREATE_AT")

    class Meta:
        db_table = 'LOG'
