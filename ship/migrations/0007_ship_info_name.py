# Generated by Django 3.1.7 on 2021-04-02 14:50

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('ship', '0006_auto_20210323_2301'),
    ]

    operations = [
        migrations.AddField(
            model_name='ship',
            name='info_name',
            field=models.CharField(blank=True, db_index=True, max_length=100, null=True),
        ),
    ]