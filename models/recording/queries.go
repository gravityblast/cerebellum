package recording

const queryAllByReleaseGid = `
  SELECT
    REC.gid, REC.name, REC.comment, REC.length
  FROM
    recording REC
  JOIN  track T
    ON T.recording = REC.id
  JOIN medium M
    ON T.medium = M.id
  JOIN release REL
    ON M.release = REL.id
  WHERE
    REL.gid = $1
  ORDER BY M.position, T.position;`

const queryByGid = `
  SELECT
    R.gid, R.name, R.comment, R.length, R.artist_credit
  FROM
    recording R
  WHERE
    R.gid = $1 limit 1;`

const queryByReleaseGidAndGid = `
  SELECT
    REC.gid, REC.name, REC.comment, REC.length, REC.artist_credit
  FROM
    recording REC
  JOIN track T
    ON REC.id = T.recording
  JOIN medium M
    ON T.medium = M.id
  JOIN release REL
    ON M.release = REL.id
  WHERE
    REL.gid = $1 AND
    REC.gid = $2 limit 1;`

