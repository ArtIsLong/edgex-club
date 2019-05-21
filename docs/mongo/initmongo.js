(function(){
  db=db.getSiblingDB('admin');
  try{
    db.createUser({user:'root',pwd:'edgex-club-root',roles:[{role:'userAdminAnyDatabase',db:'admin'},{role:'readWrite',db:'admin'}]});
  }catch(e){
    return ;
  }
  db.auth('root','edgex-club-root');

  db=db.getSiblingDB('edgex-club');
  db.createUser({user:'edgex-club-db',pwd:'Yanhua@1989',roles:[{role:'readWrite',db:'edgex-club'}]});
  db.auth('edgex-club-db','Yanhua@1989');
}());
