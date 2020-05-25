# BUILD:  docker build -t my-perl-app .
# RUN: docker run -it --rm --name my-running-app my-perl-app
FROM perl:5.20
RUN curl -L http://cpanmin.us | perl - App::cpanminus
RUN cpanm Data::RandomPerson
COPY . /usr/src/myapp
WORKDIR /usr/src/myapp
CMD [ "perl", "./echoit.pl" ]
