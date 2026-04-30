from django.http import HttpResponse
from django.shortcuts import render

def home_index(request):
    return render(request, 'home.html')
def about_index(request):
    return render(request, 'about.html')
def contact_index(request):
    return render(request, 'contact.html')