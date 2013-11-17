package release

const queryByGid = `
  SELECT
    R.gid, R.name, R.comment, R.artist_credit, RS.name, RP.name, RGPT.name as type
  FROM
    release R
  JOIN release_group RG
    ON R.release_group = RG.id
  LEFT JOIN release_group_primary_type RGPT
    ON RG.type = RGPT.id
  LEFT JOIN release_status RS
    ON R.status = RS.id
  LEFT JOIN release_packaging RP
    ON R.packaging = RP.id
  WHERE
    R.gid = $1 limit 1;`

const queryExists = `SELECT 1 FROM release where gid = $1`

const queryAllByArtistGid = `
  SELECT
    gid, name, comment, status, type, packaging, date_year, date_month, date_day
  FROM (
    SELECT
      row_number() OVER(PARTITION BY R.release_group ORDER BY RC.date_year, RC.date_month, RC.date_day) as release_row_number,
      R.gid, R.name, R.comment, RS.name as status, RGPT.name as type, RP.name as packaging, RC.date_year, RC.date_month, RC.date_day
    FROM
      release R
    JOIN release_group RG
      ON R.release_group = RG.id
    JOIN artist_credit_name ACN
      ON R.artist_credit = ACN.artist_credit
    JOIN artist A
      ON ACN.artist = A.id
    LEFT JOIN (
      SELECT release, country, date_year, date_month, date_day FROM release_country
      UNION ALL
      SELECT release, NULL, date_year, date_month, date_day FROM release_unknown_country
    ) RC
      ON RC.release = R.id
    LEFT JOIN release_status RS
      ON R.status = RS.id
    LEFT JOIN release_group_primary_type RGPT
      ON RG.type = RGPT.id
    LEFT JOIN release_packaging RP
      ON R.packaging = RP.id
    WHERE
      A.gid = $1
    ORDER BY
      release_row_number, RC.date_year, RC.date_month, RC.date_day
  ) RELEASES
  WHERE RELEASES.release_row_number = 1;`

const queryAllByReleaseGroupGid = `
  SELECT
    R.gid, R.name, R.comment, RS.name as status, RGPT.name as type, RP.name as packaging
  FROM
    release R
  JOIN release_group RG
    ON R.release_group = RG.id
  LEFT JOIN release_status RS
    ON R.status = RS.id
  LEFT JOIN release_group_primary_type RGPT
    ON RG.type = RGPT.id
  LEFT JOIN release_packaging RP
    ON R.packaging = RP.id
  WHERE
    RG.gid = $1;`
