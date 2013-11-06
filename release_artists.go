package main

const FindReleaseArtistsByArtistCreditQuery = `
  SELECT
    A.gid, A.name
  FROM
    artist_credit_name ACN
  JOIN artist A
    on A.id = ACN.artist
  WHERE
    ACN.artist_credit = $1;`

func FindReleaseArtistsByArtistCredit(artistCredit int) []*ReleaseArtist {
  artists := make([]*ReleaseArtist, 0)

  rows, err := DB.Query(FindReleaseArtistsByArtistCreditQuery, artistCredit)
  if err != nil {
    return artists
  }

  for rows.Next() {
    artist := &ReleaseArtist{}
    err := rows.Scan(&artist.Gid, &artist.Name)
    if err == nil {
      artists = append(artists, artist)
    }
  }

  return artists
}
