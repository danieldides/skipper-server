CREATE TABLE ads (
	id SERIAL PRIMARY KEY UNIQUE NOT NULL,
	video_id TEXT,
	beginning INT,
	duration INT,
	score INT
);

INSERT INTO public.ads
(video_id, beginning, duration, score)
VALUES('1H-HOPiai5Y', 10, 25, 5);

INSERT INTO public.ads
(video_id, beginning, duration, score)
VALUES('1H-HOPiai5Y', 10, 20, -5);
