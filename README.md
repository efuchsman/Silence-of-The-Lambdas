<h1 align="center"> Silence of the Lambdas</h1>

 <div align="center">
  <a href="https://github.com/efuchsman/Silence-of-The-Lambdas">
    <img src = "https://media.giphy.com/media/v1.Y2lkPTc5MGI3NjExbGVjcXJpbGJhNHcyYXNiMHZhNTA5anY1ZGZhcjNhdTNtd2pudWs5eCZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/l2SpUEkYKiyRewWqc/giphy.gif">
  </a>
</div>

  <h3 align="center">
    <br />
  </h3>
</div>

# Table of Contents
* [App Description](#app-description)
* [Learning Goals](#learning-goals)
* [System Requirements](#system-requirements)
* [Technologies Used](#technologies-used)
* [Respository](#repository)
* [Endpoints](#endpoints)
* [Resources](#resources)
* [Contact](#contact)

# App Description:
* Silence of the Lambdas allows users to search for the killers from the movies, and the victims of those killers.

# Learning Goals
* Build a Lambda function that is uses resource triggers from the API Gateway
* Integrate with DynamoDB
* Provision deployments with Terraform
* Provision building and testing with Bazel

# System Requirements
* Darwin AMD64
* Go 1.21
* Terraform v1.5.7
* Bazel 7.0.2

# Technologies Used
[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![AWS Lambda](https://img.shields.io/badge/AWS_Lambda-FF9900?style=for-the-badge&logo=amazon-aws&logoColor=white)](https://aws.amazon.com/lambda/)
[![AWS API Gateway](https://img.shields.io/badge/AWS_API_Gateway-FF9900?style=for-the-badge&logo=amazon-aws&logoColor=white)](https://aws.amazon.com/api-gateway/)
[![DynamoDB](https://img.shields.io/badge/DynamoDB-4053D6?style=for-the-badge&logo=amazon-dynamodb&logoColor=white)](https://aws.amazon.com/dynamodb/)
[![Bazel](https://img.shields.io/badge/Bazel-white?style=for-the-badge&logo=data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCA1MTIgNTEyIiB3aWR0aD0iNTEyIiBoZWlnaHQ9IjUxMiI+CiAgPHBhdGggc3R5bGU9ImZpbGw6Izc2RDI3NTsiICAgICAgZD0iTTE0NCAzMiBsMTEyIDExMiAtMTEyIDExMiBsLTExMiAtMTEyeiIvPgogIDxwYXRoIHN0eWxlPSJmaWxsOiM0M0EwNDc7IiAgICBkPSJNMzIgMTQ0IHYxMTIgbDExMiAxMTIgdi0xMTJ6Ii8+CgogIDxwYXRoIHN0eWxlPSJmaWxsOiM3NkQyNzU7IiAgICAgIGQ9Ik0zNjggMzIgIGwxMTIgMTEyIC0xMTIgMTEyIC0xMTIgLTExMnoiLz4KICA8cGF0aCBzdHlsZT0iZmlsbDojNDNBMDQ3OyIgICAgIGQ9Ik00ODAgMTQ0IHYxMTIgbC0xMTIgMTEyIHYtMTEyeiIvPgoKICA8cGF0aCBzdHlsZT0iZmlsbDojNDNBMDQ3OyIgICAgIGQ9Ik0yNTYgMTQ0IGwxMTIgMTEyIC0xMTIgMTEyIC0xMTIgLTExMnoiLz4KICA8cGF0aCBzdHlsZT0iZmlsbDojMDA3MDFBOyIgIGQ9Ik0yNTYgMzY4IHYxMTIgbC0xMTIgLTExMiAgdi0xMTJ6Ii8+CiAgPHBhdGggc3R5bGU9ImZpbGw6IzAwNDMwMDsiIGQ9Ik0yNTYgMzY4IGwxMTIgLTExMiB2MTEyIGwtMTEyIDExMnoiLz4KPC9zdmc+Cg==)](https://bazel.build/)
[![Terraform](https://img.shields.io/badge/Terraform-623CE4?style=for-the-badge&logo=terraform&logoColor=white)](https://www.terraform.io/)

# Repository
* https://github.com/efuchsman/Silence-of-the-Lambdas

# Endpoints

## Get a Killer
**Note:** Due to time constraints the only available Killers are HannibalLecter and JameGumb.

`GET 'https://795z14o07b.execute-api.us-west-2.amazonaws.com/beta/killers/{full_name}'`

```
GET https://795z14o07b.execute-api.us-west-2.amazonaws.com/beta/killers/HannibalLecter

{
  "full_name": "HannibalLecter",
  "first_name": "Hannibal",
  "last_name": "Lecter",
  "movie_actors": [
    "Anthony Hopkins",
    "Brian Cox",
    "Gaspard Ulliel"
  ],
  "movies": [
    "Hannibal",
    "Hannibal Rising",
    "Manhunter",
    "Red Dragon",
    "The Silence of the Lambs"
  ],
  "nickname": "Hannibal the Cannibal",
  "profession": "Forensic Psychiatrist",
  "image": "https://static.wikia.nocookie.net/hannibal/images/2/23/Voonakeste-vaikimine-73591313.jpg/revision/latest?cb=20210507145148"
}
```

## Get Victims for a Killer
**Note:** Due to time constraints the lists of victims are abbreviated.

`GET 'https://795z14o07b.execute-api.us-west-2.amazonaws.com/beta/killers/{full_name}/victims'`

```
GET https://795z14o07b.execute-api.us-west-2.amazonaws.com/beta/killers/HannibalLecter/victims

{
  "Victims": [
    {
      "full_name": "EnrikasDortlich",
      "first_name": "Enrikas",
      "last_name": "Dortlich",
      "actor": "Richard Brake",
      "movie": "Hannibal Rising",
      "cause_of_death": "Beheaded with rope tied to a horse",
      "occupation": "SS Officer in Border Guards",
      "cannibalized": true,
      "image": "https://static.wikia.nocookie.net/hannibal/images/6/6d/Dortlich5.png/revision/latest?cb=20140407222947"
    },
    {
      "full_name": "PaulMomund",
      "first_name": "Paul",
      "last_name": "Momund",
      "actor": "Charles Maquignon",
      "movie": "Hannibal Rising",
      "cause_of_death": "Beheaded with a tantō after slashing across the stomach",
      "occupation": "Butcher",
      "cannibalized": false,
      "image": "https://static.wikia.nocookie.net/hannibal/images/c/c1/Paul1.png/revision/latest?cb=20140410224238"
    },
    {
    "full_name": "ZigmasMilko",
    "first_name": "Zigmas",
    "last_name": "Milko",
    "actor": "Stephen Walters",
    "movie": "Hannibal Rising",
    "cause_of_death": "Drowned in a vat of formaldehyde",
    "occupation": "Hitman",
    "cannibalized": true,
    "image": "https://static.wikia.nocookie.net/hannibal/images/2/25/Milko5.png/revision/latest?cb=20140411212533"
    }
  ]
}

```

# Resources

- https://hannibal.fandom.com/wiki/Hannibal%27s_Kill_List

# Contact

<table align="center">
  <tr>
    <td><img src="https://avatars.githubusercontent.com/u/104859844?s=150&v=4"></td>
  </tr>
  <tr>
    <td>Eli Fuchsman</td>
  </tr>
  <tr>
    <td>
      <a href="https://github.com/efuchsman">GitHub</a><br>
      <a href="https://www.linkedin.com/in/elifuchsman/">LinkedIn</a>
   </td>
  </tr>
</table>
