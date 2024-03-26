-- @block
CREATE TABLE Users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    pw VARCHAR(255) NOT NULL
);

-- @block
INSERT INTO Users (username, pw)
VALUES 
(
    'MillingRoom', 'cnc123'
    ), 
('PaintRoom', 'spraypaint123'),
('GlazingRoom', 'windows123'),
('ShippingAndReceiving', 'forklift123'),
('Admin', 'admin123');

-- @block
SELECT * FROM Users;