package release

const queryByGid = `
  SELECT
    R.gid, R.name, R.comment, R.artist_credit, RS.name, RP.name
  FROM
    release R
  LEFT JOIN release_status RS
    ON R.status = RS.id
  LEFT JOIN release_packaging RP
    ON R.packaging = RP.id
  WHERE
    R.gid = $1 limit 1;`

