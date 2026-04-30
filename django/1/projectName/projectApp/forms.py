from django import forms

class URLForm(forms.form):
    url = forms.URLForm(
        label = 'Enter URL',
        widget = forms.TextInput(attrs={
            'placeholder': 'example.com',
        })
    )