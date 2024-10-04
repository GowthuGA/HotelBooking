CREATE TABLE hotels (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    address VARCHAR(255) NOT NULL,
    city VARCHAR(100) NOT NULL,
    country VARCHAR(60) NOT NULL,
    rating DECIMAL(2, 1),  -- Rating out of 5.0
    phone VARCHAR(15),
    email VARCHAR(100),
    is_active BOOLEAN DEFAULT TRUE,
    CONSTRAINT hotels_pk PRIMARY KEY (id)
) ENGINE = InnoDB;

CREATE INDEX hotels_city ON hotels (city);
CREATE INDEX hotels_country ON hotels (country);

-- rooms
CREATE TABLE rooms (
    id INT NOT NULL AUTO_INCREMENT,
    hotel_id INT NOT NULL,
    room_number VARCHAR(10) NOT NULL,
    room_type VARCHAR(50) NOT NULL,  -- e.g., Single, Double, Suite
    price DECIMAL(10, 2) NOT NULL,
    capacity INT NOT NULL,  -- Maximum number of guests
    is_available BOOLEAN DEFAULT TRUE,
    CONSTRAINT rooms_pk PRIMARY KEY (id),
    CONSTRAINT fk_rooms_hotel_id FOREIGN KEY (hotel_id)
        REFERENCES hotels (id)
        ON UPDATE NO ACTION
        ON DELETE CASCADE
) ENGINE = InnoDB;

CREATE INDEX rooms_hotel_id ON rooms (hotel_id);
CREATE INDEX rooms_type ON rooms (room_type);

-- customers
CREATE TABLE customers (
    id INT NOT NULL AUTO_INCREMENT,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    phone VARCHAR(15),
    CONSTRAINT customers_pk PRIMARY KEY (id)
) ENGINE = InnoDB;

-- admins
CREATE TABLE admins (
    id INT NOT NULL AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,  -- Store hashed passwords
    email VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_login TIMESTAMP NULL,
    CONSTRAINT admins_pk PRIMARY KEY (id)
) ENGINE = InnoDB;

-- bookings
CREATE TABLE bookings (
    id INT NOT NULL AUTO_INCREMENT,
    customer_id INT NOT NULL,
    room_id INT NOT NULL,
    hotel_id INT NOT NULL,  -- Added hotel_id for hotel information
    check_in DATE NOT NULL,
    check_out DATE NOT NULL,
    booking_date DATE NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    status ENUM('Pending', 'Confirmed', 'Cancelled') DEFAULT 'Pending',
    CONSTRAINT bookings_pk PRIMARY KEY (id),
    CONSTRAINT fk_bookings_customer_id FOREIGN KEY (customer_id)
        REFERENCES customers (id)
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_bookings_room_id FOREIGN KEY (room_id)
        REFERENCES rooms (id)
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_bookings_hotel_id FOREIGN KEY (hotel_id)
        REFERENCES hotels (id)
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
) ENGINE = InnoDB;

CREATE INDEX bookings_customer_id ON bookings (customer_id);
CREATE INDEX bookings_room_id ON bookings (room_id);
CREATE INDEX bookings_hotel_id ON bookings (hotel_id);  -- Index for hotel_id
CREATE INDEX bookings_check_in ON bookings (check_in);
