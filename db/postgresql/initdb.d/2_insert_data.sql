-- users
INSERT INTO users ("id","name")
VALUES ('b9d4a4ab-ea45-d22f-3ed6-46c32ec8b2b1','mahiro'),
        ('158912d3-5a99-6fa6-1468-c5a8e28693c0','doer');

-- todos
INSERT INTO todos ("id","user_id","name")
VALUES ('e1d388da-47e0-ca1d-2722-0a41d93f872f','b9d4a4ab-ea45-d22f-3ed6-46c32ec8b2b1','task1'),
        ('91d1a74f-8d0e-7287-b218-31e8fdaf8ad5','b9d4a4ab-ea45-d22f-3ed6-46c32ec8b2b1','task2'),
        ('6b72971a-7245-2bbd-ec5d-c38045722c2f','b9d4a4ab-ea45-d22f-3ed6-46c32ec8b2b1','task3'),
        ('57d4950b-2764-f729-0cf1-061018bf0718','b9d4a4ab-ea45-d22f-3ed6-46c32ec8b2b1','task4'),
        ('1ad482bb-23a1-9dca-1672-cd1dca55342d','b9d4a4ab-ea45-d22f-3ed6-46c32ec8b2b1','task5');