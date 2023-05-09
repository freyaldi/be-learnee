-- Insert data into the users table
INSERT INTO users (email, password, is_admin, fullname, address, phone_number, referral) VALUES
	('kim.minju@gmail.com', 'password012', true, 'Kim Minju', 'Jl. Gwangju No. 33, Seoul', '08222233344', 'REFER012');

-- Insert data into the categories table
INSERT INTO categories (name) VALUES
	('Machine Learning'),
	('Computer Vision'),
	('Natural Language Processing'),
	('Robotics');

-- Insert data into the tags table
INSERT INTO tags (name) VALUES
	('Python'),
	('TensorFlow'),
	('PyTorch'),
	('Deep Learning'),
	('Neural Networks'),
	('Image Recognition'),
	('Speech Recognition'),
	('Chatbots'); 

-- Insert data into the courses table
INSERT INTO courses (title, slug, summary_description, content, img_thumbnail, img_url, author_name, category_id, tag_id, price, created_at, updated_at) VALUES
	('Introduction to Machine Learning', 'intro-to-ml', 'Learn the basics of machine learning', 'In this course, you will learn the basics of machine learning, including supervised and unsupervised learning, and how to use Python and TensorFlow to build machine learning models.', 'https://example.com/thumbnails/ml.png', 'https://example.com/images/ml.png', 'Mari M', 1, 2, 149000, NOW(), NOW()),
	('Computer Vision Fundamentals', 'cv-fundamentals', 'Learn the fundamentals of computer vision', 'In this course, you will learn the fundamentals of computer vision, including image processing and object recognition, and how to use OpenCV and Python to build computer vision models.', 'https://example.com/thumbnails/cv.png', 'https://example.com/images/cv.png', 'Shinji I', 2, 5, 119000, NOW(), NOW()),
	('Natural Language Processing with PyTorch', 'nlp-with-pytorch', 'Learn how to process and analyze natural language data', 'In this course, you will learn how to process and analyze natural language data, including text classification and sentiment analysis, using PyTorch and Python.', 'https://example.com/thumbnails/nlp.png', 'https://example.com/images/nlp.png', 'Rei A', 3, 3, 99000, NOW(), NOW()),
	('Robotics Programming with ROS', 'ros-programming', 'Learn how to program robots using ROS', 'In this course, you will learn how to program robots using the Robot Operating System (ROS), including robot kinematics and navigation, and how to use Python to control robots.', 'https://example.com/thumbnails/ros.png', 'https://example.com/images/ros.png', 'Kaworu N', 4, 8, 249000, NOW(), NOW());

-- Insert data into the gifts table
INSERT INTO gifts (name, stock) VALUES
	('Vitamin', 50),
	('T-Shirt', 20),
	('Mug', 10),
	('Keychain', 20);

-- Insert data into the vouchers table
INSERT INTO vouchers (name, voucher_code, benefit) VALUES
    ('Winter Sale', 'WINTER21', 20.0),
    ('Summer Sale', 'SUMMER22', 25.0);

