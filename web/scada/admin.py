from django.contrib import admin
from django.contrib.auth.models import User, Group
from .models import Device, Log


class DeviceAdmin(admin.ModelAdmin):
    fields = ('mac_id', 'name', 'host', 'port', 'unit_id', 'remap_version', 'process_interval', 'retry_cycle', 'retry_count', 'retry_conn_failed_count', 'enabled',)
    list_display = ('id', 'mac_id', 'name', 'host', 'port', 'unit_id', 'enabled',)
    actions = ('change_enabled_devices', 'change_disabled_devices',)
    search_fields = ('mac_id', 'name',)

    def change_enabled_devices(self, request, queryset):
        update_count = queryset.update(enabled=True)
        self.message_user(request, f'{update_count}건의 장치를 enabled 상태로 변경합니다.')
    change_enabled_devices.short_description = '선택된 장치들을 enabled 상태로 변경합니다.'

    def change_disabled_devices(self, request, queryset):
        update_count = queryset.update(enabled=False)
        self.message_user(request, f'{update_count}건의 장치를 disabled 상태로 변경합니다.')
    change_disabled_devices.short_description = '선택된 장치들을 disabled 상태로 변경합니다.'


class LogAdmin(admin.ModelAdmin):
    readonly_fields = ('timestamp', 'mac_id', 'name', 'log',)
    list_display = ('timestamp', 'mac_id', 'name', 'log',)
    ordering = ('-create_at',)
    search_fields = ('log',)

    def timestamp(self, obj):
        return obj.create_at.strftime('%Y-%m-%d %H:%M:%S.%f')
    timestamp.short_description = 'Log create_at'

    def mac_id(self, obj):
        return Device.objects.get(id=obj.fk_id).mac_id
    mac_id.short_description = 'Mac ID'

    def name(self, obj):
        return Device.objects.get(id=obj.fk_id).name

    def has_add_permission(self, request):
        return False


admin.site.register(Device, DeviceAdmin)
admin.site.register(Log, LogAdmin)

admin.site.unregister(User)
admin.site.unregister(Group)
