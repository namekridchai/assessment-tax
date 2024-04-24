
CREATE TABLE IF NOT EXISTS public.allowance_master
(
    personal FLOAT NOT NULL,
    k_receipt FLOAT  NOT NULL
);

INSERT INTO public.allowance_master (personal,k_receipt)
VALUES (50000,50000);