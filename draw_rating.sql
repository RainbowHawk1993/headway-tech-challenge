-- DrawRating
   -----------
-- Converts a numeric score (0 – 100) into a five-star string.

FUNCTION DrawRating(score):

    IF score >= 0 AND score <= 20 THEN
        RETURN "★☆☆☆☆"

    ELSE IF score > 20 AND score <= 40 THEN
        RETURN "★★☆☆☆"

    ELSE IF score > 40 AND score <= 60 THEN
        RETURN "★★★☆☆"

    ELSE IF score > 60 AND score <= 80 THEN
        RETURN "★★★★☆"

    ELSE IF score > 80 AND score <= 100 THEN
        RETURN "★★★★★"

    ELSE
        RETURN "☆☆☆☆☆"

END FUNCTION
