package releasegroup

const queryAllByArtistGid = `
  SELECT
    RG.gid, RG.name, RG.comment, RG.artist_credit, RGPT.name,
    RGM.first_release_date_year, RGM.first_release_date_month,
    RGM.first_release_date_day
  FROM
    release_group RG
  JOIN artist_credit_name ACN
    ON RG.artist_credit = ACN.artist_credit
  JOIN artist A
    ON ACN.artist = A.id
  LEFT JOIN release_group_primary_type RGPT
    ON RG.type = RGPT.id
  LEFT JOIN release_group_meta RGM
    ON RG.id = RGM.id
  WHERE
    A.gid = $1
  ORDER BY
    RGM.first_release_date_year, RGM.first_release_date_month, RGM.first_release_date_day;`

const queryByGid = `
  SELECT
    RG.gid, RG.name, RG.comment, RG.artist_credit, RGPT.name,
    RGM.first_release_date_year, RGM.first_release_date_month,
    RGM.first_release_date_day
  FROM
    release_group RG
  LEFT JOIN release_group_primary_type RGPT
    ON RG.type = RGPT.id
  LEFT JOIN release_group_meta RGM
    ON RG.id = RGM.id
  WHERE
    RG.gid = $1 limit 1;`

const queryByArtistGidAndGid = `
  SELECT
    RG.gid, RG.name, RG.comment, RG.artist_credit, RGPT.name,
    RGM.first_release_date_year, RGM.first_release_date_month,
    RGM.first_release_date_day
  FROM
    release_group RG
  JOIN artist_credit_name ACN
    ON RG.artist_credit = ACN.artist_credit
  JOIN artist A
    ON ACN.artist = A.id
  LEFT JOIN release_group_primary_type RGPT
    ON RG.type = RGPT.id
  LEFT JOIN release_group_meta RGM
    ON RG.id = RGM.id
  WHERE
    A.gid = $1 AND
    RG.gid = $2 limit 1;`

const queryExists = `SELECT 1 FROM release_group where gid = $1`
