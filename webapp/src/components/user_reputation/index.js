import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';

import {id as pluginId} from '../../manifest';
import {fetchUserReputations} from '../../actions';
import {getUserReputations} from '../../selectors';

import UserReputation from './user_reputation.jsx';

const REDUCER = `plugins-${pluginId}`;

function mapStateToProps(state, ownProps) {
    const userId = ownProps.user ? ownProps.user.id : '';
    const userReputations = getUserReputations(state)[userId] || null;
    return {
        userId,
        userReputations,
    };
}

function mapDispatchToProps(dispatch) {
    return {
        actions: bindActionCreators({
            fetchUserReputations,
        }, dispatch),
    };
}

export default connect(mapStateToProps, mapDispatchToProps)(UserReputation);