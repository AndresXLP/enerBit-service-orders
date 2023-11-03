CREATE TYPE STATUS as enum ('new', 'done','cancelled');

CREATE TABLE IF NOT EXISTS public.work_orders(

    id                 UUID      NOT NULL PRIMARY KEY,
    customer_id        UUID      NOT NULL,
    title              TEXT      NOT NULL,
    planned_date_begin TIMESTAMP NOT NULL,
    planned_date_end   TIMESTAMP NOT NULL,
    status             STATUS    NOT NULL,
    created_at         TIMESTAMP NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES customers (id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT
);

INSERT INTO public.work_orders (id, customer_id, title, planned_date_begin, planned_date_end, status, created_at)
VALUES
    ('246e642d-e3bf-4ba3-96e7-e88d04633653','508c9897-30c7-47f0-83ed-f1fbf60ff73b', 'Instalacion Medidor', '2023-11-01 08:00:00', '2023-11-01 08:45:00', 'new',now()),
    ('7057153d-796b-4520-951a-52e58f77224e','d17ef45b-19be-41b9-b824-fad49cf0a28c', 'Revision Medidor', '2023-11-02 14:00:00', '2023-11-02 15:45:00', 'new',now()),
    ('50bf25fb-4f0e-44a5-96e5-9bfafffac50d','f502f315-6c20-4028-bdae-f4a77f852ce9', 'Cambio Medidor', '2023-11-03 10:00:00', '2023-11-03 12:00:00', 'new',now()),
    ('721ec232-a858-43ba-95a1-370f0f90b5de','dd265a87-d33f-4757-82d6-72e34198b3e3', 'Instalacion Medidor', '2023-10-31 10:00:00', '2023-10-31 12:00:00', 'done',now()),
    ('9bda51bd-7443-4a97-9510-e76f64faf732','80796cdf-e7d9-455b-953d-27ee82c97085', 'Instalacion Medidor', '2023-10-25 10:00:00', '2023-10-25 12:00:00', 'done',now()),
    ('31d3dfe5-eb4b-4a4b-988e-a515de392b76','80796cdf-e7d9-455b-953d-27ee82c97085', 'Cambio Medidor', '2023-10-28 10:00:00', '2023-10-28 12:00:00', 'new',now()),
    ('fdc24871-5db4-4d75-95b8-eb541c72f47c','4022206f-fee1-4810-b263-42560fa40f90', 'Instalacion Medidor', '2023-10-22 10:00:00', '2023-10-22 12:00:00', 'done',now()),
    ('1b7a2331-afe0-4a0b-9c57-9a5696d1909a','75e33974-42d6-4041-b861-65dc54e1db79', 'Instalacion Medidor', '2023-10-23 10:00:00', '2023-10-23 12:00:00', 'done',now()),
    ('09af95e6-5018-45fd-a00c-b2cefb9e1dec','1f5ef660-e441-4018-aba2-d081866a0e68', 'Instalacion Medidor', '2023-10-24 10:00:00', '2023-10-24 12:00:00', 'done',now());






