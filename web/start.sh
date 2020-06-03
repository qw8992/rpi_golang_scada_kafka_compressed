# parameters
WEB_ADMIN_EMAIL=${WEB_ADMIN_EMAIL:-"admin@itsroom.com"}
WEB_ADMIN_PASSWORD=${WEB_ADMIN_PASSWORD:-"admin"}

echo Start Makemigrations 
python3 manage.py makemigrations --noinput 

echo Start Migrate 
python3 manage.py migrate --noinput 

echo Create Admin User
echo "from django.contrib.auth.models import User; ""User.objects.create_superuser('admin', '$WEB_ADMIN_EMAIL', '$WEB_ADMIN_PASSWORD') if not User.objects.filter(username='admin') else False;" | python3 manage.py shell

echo Start Server 
python3 manage.py runserver 0.0.0.0:8000 
