import React, {Component} from 'react'
import Swiper from 'react-native-deck-swiper'
import {Button, StyleSheet, Text, View, Image} from 'react-native'
import externalStyles from './src/pages/styles.js'

export default class App extends Component {
  constructor (props) {
    super(props)
    this.state = { cards: [{'uri':'./src/pages/user.png','fname':'Jamie','lname':'Franco','decription':'I love metal!'},{'uri':'./src/pages/user.png','fname':'a','lname':'a','decription':'a'}],
            swipedAllCards: false,
            swipeDirection: '',
            isSwipingBack: false,
            cardIndex: 0 }
  }

  componentDidMount() {
    fetch('http://138.16.51.208:3000/getRemainingUsersForEvent', {
      method: 'POST',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        user_id: "16",
        event_id: "1"
      }),
    })
    .then(response => response.json())
    .then(data => this.setState({ cards: data,
                    swipedAllCards: false,
                    swipeDirection: '',
                    isSwipingBack: false,
                    cardIndex: 0 }))
    .catch((error)=>{
       console.log("Api call error");
       alert(error.message);
    });
    
  }

  renderCard = card => {
    return (
      <View style={styles.card}>
        <Image source={require('./src/pages/user.png')} style={{marginTop: 50, height:200, width:200, borderRadius:100, justifyContent: 'center', alignItems: 'center', marginLeft: 65}}/>
        <Text style={styles.nameStyle}>{card.fname} {card.lname}</Text>
        <Text style={styles.descriptionStyle}>{card.description}</Text>
      </View>
    )
  };

  onSwipedAllCards = () => {
    this.setState({
      swipedAllCards: true
    })
  };

  swipeBack = () => {
    if (!this.state.isSwipingBack) {
      this.setIsSwipingBack(true, () => {
        this.swiper.swipeBack(() => {
          this.setIsSwipingBack(false)
        })
      })
    }
  };

  setIsSwipingBack = (isSwipingBack, cb) => {
    this.setState(
      {
        isSwipingBack: isSwipingBack
      },
      cb
    )
  };

  swipeLeft = () => {
    this.swiper.swipeLeft()
  };

  render () {
    return (
      <View style={styles.container}>
        
        <Swiper
          ref={swiper => {
            this.swiper = swiper
          }}
          onSwiped={this.onSwiped}
          onTapCard={this.swipeLeft}
          cards={this.state.cards}
          cardIndex={this.state.cardIndex}
          cardVerticalMargin={80}
          renderCard={this.renderCard}
          onSwipedAll={this.onSwipedAllCards}
          overlayLabels={{
            bottom: {
              title: 'BLEAH',
              style: {
                label: {
                  backgroundColor: 'black',
                  borderColor: 'black',
                  color: 'white',
                  borderWidth: 1
                },
                wrapper: {
                  flexDirection: 'column',
                  alignItems: 'center',
                  justifyContent: 'center'
                }
              }
            },
            left: {
              title: 'NOPE',
              style: {
                label: {
                  backgroundColor: 'black',
                  borderColor: 'black',
                  color: 'white',
                  borderWidth: 1
                },
                wrapper: {
                  flexDirection: 'column',
                  alignItems: 'flex-end',
                  justifyContent: 'flex-start',
                  marginTop: 30,
                  marginLeft: -30
                }
              }
            },
            right: {
              title: 'LIKE',
              style: {
                label: {
                  backgroundColor: 'black',
                  borderColor: 'black',
                  color: 'white',
                  borderWidth: 1
                },
                wrapper: {
                  flexDirection: 'column',
                  alignItems: 'flex-start',
                  justifyContent: 'flex-start',
                  marginTop: 30,
                  marginLeft: 30
                }
              }
            },
            top: {
              title: 'SUPER LIKE',
              style: {
                label: {
                  backgroundColor: 'black',
                  borderColor: 'black',
                  color: 'white',
                  borderWidth: 1
                },
                wrapper: {
                  flexDirection: 'column',
                  alignItems: 'center',
                  justifyContent: 'center'
                }
              }
            }
          }}
          animateOverlayLabelsOpacity
          animateCardOpacity
        >
          <Button onPress={this.swipeLeft} title='Swipe Left' />
        </Swiper>
      </View>
    )
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#F5FCFF'
  },
  card: {
    flex: 1,
    borderRadius: 4,
    borderWidth: 2,
    borderColor: '#E8E8E8',
    justifyContent: 'center',
    backgroundColor: 'white'
  },
  nameStyle: {
    flex: 2,
    textAlign: 'center',
    fontSize: 35,
    backgroundColor: 'transparent',
    marginTop: 25
  },
  descriptionStyle: {
    flex: 2,
    marginLeft: 10,
    marginRight: 10,
    fontSize: 20,
    backgroundColor: 'transparent',
    marginTop: 0
  },
  done: {
    textAlign: 'center',
    fontSize: 30,
    color: 'white',
    backgroundColor: 'transparent'
  }
})
