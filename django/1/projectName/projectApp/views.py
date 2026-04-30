from django.shortcuts import render, redirect
from django.http import HttpResponse
from .forms import URLForm
from .models import ValidateURL

# Create your views here.
def index(request):
    if request.method == 'POST':
        form = URLForm(request.POST)
        if form.is_valid():
            url = form.cleaned_data['url']
            ValidateURL.objects.create(url = url)
            return redirect('success')
        else:
            form = URLForm()
        return render(request, 'templates/home.html', {
            'form' : form
        })
def success(request):
    validated_urls = ValidateURL.objects.all()
    return render(request, 'templates/contact.html', {
        'validated_urls' : validated_urls
    })