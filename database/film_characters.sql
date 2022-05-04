select *
from resources_people
where id in(
    select people_id
    from resources_film_characters
    where film_id = ?
);
