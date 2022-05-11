# Go Rest Clean Architecture

Finding the best recipe for architecture

## Normalizing SQL
```
UPDATE payments SET bank_account_name='' WHERE bank_account_name IS NULL;
UPDATE payments SET user_email='' WHERE user_email IS NULL;
```