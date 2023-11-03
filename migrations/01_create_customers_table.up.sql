CREATE TABLE IF NOT EXISTS public.customers
(
    id         UUID      NOT NULL PRIMARY KEY,
    first_name TEXT      NOT NULL,
    last_name  TEXT      NOT NULL,
    address    TEXT      NOT NULL,
    start_date TIMESTAMP NULL,
    end_date   TIMESTAMP NULL,
    is_active  BOOLEAN   NOT NULL,
    created_at TIMESTAMP NOT NULL
);

INSERT INTO public.customers (id, first_name, last_name, address, start_date, end_date, is_active, created_at)
VALUES
    ('508c9897-30c7-47f0-83ed-f1fbf60ff73b', 'John', 'Doe', '123 Main St', null, null, false, now()),
    ('d17ef45b-19be-41b9-b824-fad49cf0a28c', 'Jane', 'Smith', '456 Elm St', null, null, false, now()),
    ('f502f315-6c20-4028-bdae-f4a77f852ce9', 'Michael', 'Johnson', '789 Oak St', null, null, false, now()),
    ('dd265a87-d33f-4757-82d6-72e34198b3e3', 'Emily', 'Brown', '101 Pine St', '2023-10-31 11:25:00', null, true, now()),
    ('80796cdf-e7d9-455b-953d-27ee82c97085', 'William', 'Davis', '222 Cedar St', '2023-10-25 11:50:00', '2023-10-27 08:00:00', false, now()),
    ('98fdf0cb-f6be-4acc-b67f-808e56b07fdd', 'Olivia', 'Wilson', '333 Birch St', null, null, false, now()),
    ('040a5582-ed99-4224-a910-8d24b2835f63', 'Ethan', 'Miller', '444 Maple St', null, null, false, now()),
    ('4022206f-fee1-4810-b263-42560fa40f90', 'Sophia', 'Anderson', '555 Walnut St', '2023-10-22 11:50:00', null, true, now()),
    ('75e33974-42d6-4041-b861-65dc54e1db79', 'Mason', 'Martinez', '666 Redwood St', '2023-10-23 11:50:00', null, true, now()),
    ('1f5ef660-e441-4018-aba2-d081866a0e68', 'Ava', 'Rodriguez', '777 Cedar St', '2023-10-24 11:50:00', null, true, now());
