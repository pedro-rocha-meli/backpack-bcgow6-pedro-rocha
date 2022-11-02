
DROP DATABASE IF EXISTS empresa_internet;
CREATE DATABASE empresa_internet;
USE empresa_internet;


CREATE TABLE cliente (
	dni INT NOT NULL,
	nombre VARCHAR(50),
	apellido VARCHAR(50),
	fecha_nacimiento DATETIME,
	provincia VARCHAR(50),
	ciudad VARCHAR(50),
    PRIMARY KEY (dni)
);

CREATE TABLE plan (
	id INT NOT NULL AUTO_INCREMENT,
    velocidad INT,
    precio DOUBLE,
    descuento DOUBLE,
    PRIMARY KEY (id)
);

create table cliente_plan (
	cliente_dni INT,
    plan_id INT,
    FOREIGN KEY (cliente_dni) REFERENCES cliente(dni),
    FOREIGN KEY (plan_id) REFERENCES plan(id)
);

INSERT INTO cliente (dni, nombre, apellido, fecha_nacimiento, provincia, ciudad) VALUES 
	(57438959, "Marcos", "Hanzen", "1983-03-25", "Buenos Aires", "CABA"),
    (27392127, "Rodrigo", "Baez", "2000-07-12", "Buenos Aires", "CABA"),
    (57592739, "Carlos", "Deñ Puerto", "2002-02-29", "Buenos Aires", "Torquinst"),
    (54739754, "Pablo", "Novarez", "1989-05-15", "Jujuy", "Bolson"),
    (01293493, "Manuel", "Gimenez", "2009-01-09", "La Rioja", "Safina"),
    (41738494, "Milagros", "Torres", "1994-06-08", "Mar Del Plata", "Balbanes"),
    (57397245, "Ramona", "Diaz", "1978-02-12", "Mendoza", "Lujan"),
    (12347294, "Nahuel", "Rocca", "1934-08-20", "Tucuman", "Carrio"),
    (78593477, "Tadeo", "Moralez", "2006-05-21", "Mar Del Plata", "Costa Paseo"),
    (57438795, "Jazmin", "Mancusso", "2012-09-01", "Buenos Aires", "Casanova")
;

INSERT INTO plan (velocidad, precio, descuento) VALUES
	(60, 4800.30, 30),
    (12, 300, 0),
    (70, 5000, 50),
    (23, 458, 13),
    (75, 6255, 70)
;

INSERT INTO cliente_plan(cliente_dni, plan_id) VALUES
	(57438959, 1),
    (27392127, 2),
    (57592739, 3),
    (54739754, 4),
    (01293493, 5),
    (12347294, 3),
    (78593477, 2),
    (90876345, 3),
    (98423783, 5),
    (54739754, 5),
    (90876345, 2),
    (78593477, 3),
    (09324877, 3)
;

SELECT * FROM plan;
SELECT * FROM cliente;
SELECT * FROM cliente_plan;