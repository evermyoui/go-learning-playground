from django.db import models
from django.db.models import Model

# Create your models here.
class NameModel (models.Model):
    title = models.CharField(max_length=200)
    description = models.TextField()
    status = models.BooleanField(default= False)
    # image = models.ImageField(upload_to='images/')
    def __str__(self):
        return self.title
    
class ValidateURL(models.Model):
    url = models.URLField(unique=True)
    def __str__(self):
        return self.url