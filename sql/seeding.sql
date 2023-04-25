-- Insert data into the users table
INSERT INTO users (email, password, role, fullname, address, phone_number, level, referral) VALUES
	('ferza.reyaldi@gmail.com', 'password123', 'user', 'Ferza Reyaldi', 'Jl. Merbabu No. 12, Jakarta', '08123456789', 'junior', 'REFER123'),
	('bintang.dwitama@gmail.com', 'password456', 'user', 'Bintang Dwitama', 'Jl. Semeru No. 25, Jakarta', '08234567890', 'newbie', 'REFER456'),
	('eka.rahadi@gmail.com', 'password789', 'user', 'Eka Rahadi', 'Jl. Bromo No. 9, Jakarta', '08111122233', 'master', 'REFER789'),
	('kim.minju@gmail.com', 'password012', 'admin', 'Kim Minju', 'Jl. Gwangju No. 33, Seoul', '08222233344', 'senior', 'REFER012');

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
INSERT INTO courses (title, slug, summary_description, content, img_thumbnail, img_url, author_name, category_id, tag_id) VALUES
	('Introduction to Machine Learning', 'intro-to-ml', 'Learn the basics of machine learning', 'In this course, you will learn the basics of machine learning, including supervised and unsupervised learning, and how to use Python and TensorFlow to build machine learning models.', 'https://example.com/thumbnails/ml.png', 'https://example.com/images/ml.png', 'John Smith', 1, 2),
	('Computer Vision Fundamentals', 'cv-fundamentals', 'Learn the fundamentals of computer vision', 'In this course, you will learn the fundamentals of computer vision, including image processing and object recognition, and how to use OpenCV and Python to build computer vision models.', 'https://example.com/thumbnails/cv.png', 'https://example.com/images/cv.png', 'Jane Doe', 2, 5),
	('Natural Language Processing with PyTorch', 'nlp-with-pytorch', 'Learn how to process and analyze natural language data', 'In this course, you will learn how to process and analyze natural language data, including text classification and sentiment analysis, using PyTorch and Python.', 'https://example.com/thumbnails/nlp.png', 'https://example.com/images/nlp.png', 'Bob Johnson', 3, 3),
	('Robotics Programming with ROS', 'ros-programming', 'Learn how to program robots using ROS', 'In this course, you will learn how to program robots using the Robot Operating System (ROS), including robot kinematics and navigation, and how to use Python to control robots.', 'https://example.com/thumbnails/ros.png', 'https://example.com/images/ros.png', 'Alice Lee', 4, 8);

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

