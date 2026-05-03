from django.db import models

# Create your models here.

class React(models.Model):
    name = models.CharField(max_length=20)
    detail = models.TextField()

    def __str__(self):
        return self.name