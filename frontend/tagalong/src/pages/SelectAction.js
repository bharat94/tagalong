import React from 'react';
import { StyleSheet, Text, View } from 'react-native';
import { DrawerNavigator } from 'react-navigation'

import Profile from './Profile.js'
import MyEvents from './MyEvents.js'
import SelectEvent from './SelectEvent.js'

const Navigation = DrawerNavigator(
	{
		SelectEvent: {
			screen: SelectEvent,
			navigationOptions: {
     			drawerLabel: 'Select Events',
   			}
		},
		Profile: {
			screen: Profile,
			navigationOptions: {
     			drawerLabel: 'Profile',
   			}
		},
		"My Events": {
			screen: MyEvents,
			navigationOptions: {
     			drawerLabel: 'My Events',
   			}
		}
	}
)

export default Navigation;

