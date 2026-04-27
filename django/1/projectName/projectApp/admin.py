from django.contrib import admin

# Register your models here.
from .models import NameModel

admin.site.register(NameModel)