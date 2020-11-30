CREATE TABLE public.persons (
    id Serial PRIMARY KEY,
    firstName varchar(255),
    lastName varchar(255),
    address varchar(250)
);

INSERT INTO public.persons (id, firstname, lastname, address)
VALUES
    (1, 'Luke', 'Skywalker', 'Tatooine'),
    (2, 'Leia', 'Organa', 'Alderaan'),
    (3, 'Han', 'Solo', 'Corellia');