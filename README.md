# Cerebellum

Cerebellum is a small implementation of the [Musicbrainz Web Service](https://wiki.musicbrainz.org/Development/XML_Web_Service/Version_2)
made in [Go](http://golang.org/) with the [Trafic Web Framework](http://github.com/pilu/traffic).

To compile and run Cerebellum you need Go and the Musicbrainz database.

## Installation

    go get github.com/pilu/cerebellum

Copy the config file sample to another location

    cp /go/path/src/github.com/pilu/cerebellum/traffic.conf.sample /custom/path/traffic.conf

## Usage

    TRAFFIC_CONFIG_FILE=/custom/path/traffic.conf cerebellum

## Available API

###Artist:

    /artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56

Response:

    {
      gid: "056e4f3e-d505-4dad-8ec1-d04f521cbb56",
      name: "Daft Punk",
      sortName: "Daft Punk",
      comment: "",
      beginDate: "1992",
      endDate: "",
      type: "Group"
    }

###Artist Release Groups:

    /artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56/release-groups

Response:

    [
      {
        gid: "35a4f900-60c1-3f79-8cdb-193941b78768",
        name: "The New Wave",
        comment: "",
        firstReleaseDate: "1994-04-11",
        type: "Single"
      },
      ...

###Artist Release Group:

    /artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56/release-groups/aa997ea0-2936-40bd-884d-3af8a0e064dc

Response:

    {
      gid: "aa997ea0-2936-40bd-884d-3af8a0e064dc",
      name: "Random Access Memories",
      comment: "",
      firstReleaseDate: "2013-05-17",
      type: "Album",
      artists: [
        {
          gid: "056e4f3e-d505-4dad-8ec1-d04f521cbb56",
          name: "Daft Punk"
        }
      ]
    }

###Artist Releases (first release for each release group):

    /artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56/releases

Response:

    [
      {
        gid: "95d9f3ef-c935-4a7b-b9cb-36aea57b6bae",
        name: "The New Wave",
        comment: "",
        date: "1994-04-11",
        status: "Official",
        type: "Single",
        packaging: ""
      },
      {
        gid: "b1f35e3e-0d4f-451f-93c7-1eb8809b9aa5",
        name: "Da Funk",
        comment: "",
        date: "1995",
        status: "Official",
        type: "Single",
        packaging: ""
      },
      ...

###Artist Release:

    /artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56/releases/79215cdf-4764-4dee-b0b9-fec1643df7c5

Response:

    {
      gid: "79215cdf-4764-4dee-b0b9-fec1643df7c5",
      name: "Random Access Memories",
      comment: "",
      status: "Official",
      type: "Album",
      packaging: "Jewel Case",
      artists: [
        {
          gid: "056e4f3e-d505-4dad-8ec1-d04f521cbb56",
          name: "Daft Punk"
        }
      ]
    }

###Artist Releases by Release Group:

    /artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56/release-groups/aa997ea0-2936-40bd-884d-3af8a0e064dc/releases

Response:

    [
      {
        gid: "4867ceba-ffe7-40c0-a093-45be6c03c655",
        name: "Random Access Memories",
        comment: "",
        status: "Official",
        type: "Album",
        packaging: "Cardboard/Paper Sleeve"
      },
      {
        gid: "79215cdf-4764-4dee-b0b9-fec1643df7c5",
        name: "Random Access Memories",
        comment: "",
        status: "Official",
        type: "Album",
        packaging: "Jewel Case"
      },
      ...

###Artist Release Recordings:

    /artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56/releases/79215cdf-4764-4dee-b0b9-fec1643df7c5/recordings

Response:

    [
      {
        gid: "0c871a4a-efdf-47f8-98c2-cc277f806d2f",
        name: "Give Life Back to Music",
        comment: "",
        length: 274000
      },
      {
        gid: "294a1b4d-ebc0-4b03-be25-6171d382cb58",
        name: "The Game of Love",
        comment: "",
        length: 321000
      },
      ...

###Artist Release Recording:

    /artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56/releases/79215cdf-4764-4dee-b0b9-fec1643df7c5/recordings/833f00e1-781f-4edd-90e4-e52712618862

Response:

    {
      gid: "833f00e1-781f-4edd-90e4-e52712618862",
      name: "Get Lucky",
      comment: "",
      length: 367000,
      artists: [
        {
          gid: "056e4f3e-d505-4dad-8ec1-d04f521cbb56",
          name: "Daft Punk"
        },
        {
          gid: "149f91ef-1287-46da-9a8e-87fee02f1471",
          name: "Pharrell Williams"
        }
      ]
    }

###Release Group:

    /release-groups/aa997ea0-2936-40bd-884d-3af8a0e064dc

Response:

    {
      gid: "aa997ea0-2936-40bd-884d-3af8a0e064dc",
      name: "Random Access Memories",
      comment: "",
      firstReleaseDate: "2013-05-17",
      type: "Album",
      artists: [
        {
          gid: "056e4f3e-d505-4dad-8ec1-d04f521cbb56",
          name: "Daft Punk"
        }
      ]
    }

###Release Group - Release:

    /release-groups/aa997ea0-2936-40bd-884d-3af8a0e064dc/releases

Response:

    [
      {
        gid: "4867ceba-ffe7-40c0-a093-45be6c03c655",
        name: "Random Access Memories",
        comment: "",
        status: "Official",
        type: "Album",
        packaging: "Cardboard/Paper Sleeve"
      },
      {
        gid: "79215cdf-4764-4dee-b0b9-fec1643df7c5",
        name: "Random Access Memories",
        comment: "",
        status: "Official",
        type: "Album",
        packaging: "Jewel Case"
      },
      ...


###Release:

    /releases/79215cdf-4764-4dee-b0b9-fec1643df7c5

Response:

    {
      gid: "79215cdf-4764-4dee-b0b9-fec1643df7c5",
      name: "Random Access Memories",
      comment: "",
      status: "Official",
      type: "Album",
      packaging: "Jewel Case",
      artists: [
        {
          gid: "056e4f3e-d505-4dad-8ec1-d04f521cbb56",
          name: "Daft Punk"
        }
      ]
    }

###Release Recordings:

    /releases/79215cdf-4764-4dee-b0b9-fec1643df7c5/recordings

Response:

    [
      {
        gid: "0c871a4a-efdf-47f8-98c2-cc277f806d2f",
        name: "Give Life Back to Music",
        comment: "",
        length: 274000
      },
      {
        gid: "294a1b4d-ebc0-4b03-be25-6171d382cb58",
        name: "The Game of Love",
        comment: "",
        length: 321000
      },
      ...

###Recording:

    /recordings/0c871a4a-efdf-47f8-98c2-cc277f806d2f

Response:

    {
      gid: "0c871a4a-efdf-47f8-98c2-cc277f806d2f",
      name: "Give Life Back to Music",
      comment: "",
      length: 274000,
      artists: [
        {
          gid: "056e4f3e-d505-4dad-8ec1-d04f521cbb56",
          name: "Daft Punk"
        }
      ]
    }

###Release Recording:

    /releases/79215cdf-4764-4dee-b0b9-fec1643df7c5/recordings/0c871a4a-efdf-47f8-98c2-cc277f806d2f

Response:

    {
      gid: "0c871a4a-efdf-47f8-98c2-cc277f806d2f",
      name: "Give Life Back to Music",
      comment: "",
      length: 274000,
      artists: [
        {
          gid: "056e4f3e-d505-4dad-8ec1-d04f521cbb56",
          name: "Daft Punk"
        }
      ]
    }

## Author

* [Andrea Franz](http://gravityblast.com)

