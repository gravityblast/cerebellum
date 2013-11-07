package artist

const queryByGid = `
  SELECT
    A.gid, A.name, A.sort_name, A.comment,
    A.begin_date_year, A.begin_date_month, A.begin_date_day,
    A.end_date_year, A.end_date_month, A.end_date_day, AT.name
  FROM
    artist A
  LEFT JOIN artist_type AT
    ON A.type = AT.id
  WHERE
    A.gid = $1 limit 1;`

const queryExists = `SELECT 1 FROM artist where gid = $1`

const queryAllByArtistCredit = `
  SELECT
    A.gid, A.name
  FROM
    artist_credit_name ACN
  JOIN artist A
    on A.id = ACN.artist
  WHERE
    ACN.artist_credit = $1;`

