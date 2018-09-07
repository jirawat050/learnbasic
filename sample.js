  function createNode(element) {
      return document.createElement(element);
  }

  function append(parent, el) {
    return parent.appendChild(el);
  }

  const ul = document.getElementById('authors');
  const url = 'https://gist.githubusercontent.com/DQIJAO/e14c64ea610688e70228a9fb8c649b2c/raw/6cccd444c1ef65411aa3662b112634996b837414/bnk48.json';
  fetch(url)
  .then((resp) => resp.json())
  .then(function(data) {
 
   //console.log(data.members[0].slug);
    let authors = data.members;
    return authors.map(function(author) {
     // console.log(author.slug);
      let li = createNode('li'),
          img = createNode('img'),
          span = createNode('span');
      img.src = author.profile_image;
      span.innerHTML = `${author.first_name.th} ${author.last_name.th}`;
      append(li, img);
      append(li, span);
      append(ul, li);
    })
  })
  .catch(function(error) {
     
    console.log(JSON.stringify(error));
  });   
